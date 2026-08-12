#!/usr/bin/env python3

# https://www.geeksforgeeks.org/machine-learning/text-classification-using-logistic-regression/
# Much of this was taken from or adapted
# from this tutorial

import sys
import os
import re
import random
import pandas as pd
from sklearn.feature_extraction.text import CountVectorizer, TfidfVectorizer
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
import pickle

def split_paras(para):
  chars = r'[\.\!\?]'
  sentences = re.split(chars, para)
  return sentences

def clean(line):
  line = line.strip()
  if line == "" or line.startswith("─"):
    return ""
  return line.lower()

def classify_sentence(model, vectorizer, sentence):
  sentence_vector = vectorizer.transform([sentence])
  prediction = model.predict(sentence_vector)
  if prediction[0] == 1:
    return 1
  elif prediction[0] == 0:
    return 0

script_dir      = os.path.dirname(os.path.abspath(__file__))
vectorizer_file = os.path.join(script_dir, "./vectorizer.pickle")
model_file      = os.path.join(script_dir, "./model.pickle")

#
# ONE FILE
#

if sys.argv[1] == "--file":
  file = sys.argv[2]
  if not os.path.isfile(vectorizer_file):
    print("Error: No vectorizer file.")
    sys.exit()
  if not os.path.isfile(model_file):
    print("Error: No model file.")
    sys.exit()
  vectorizer = pickle.load(open(vectorizer_file, "rb"))
  model = pickle.load(open(model_file, "rb"))
  with open(file, "r", encoding = "utf-8") as f:
    yes_count   = 0
    no_count    = 0
    f_lines     = []
    f_paras_tmp = []
    f_paras     = f.readlines()

    for para in f_paras:
      f_paras_tmp += split_paras(para)
    for line in f_paras_tmp:
      line = clean(line)
      if line != "":
        f_lines.append(line)
    for line in f_lines:
      result = classify_sentence(model, vectorizer, line)
      if result == 1:
        yes_count += 1
      elif result == 0:
        no_count += 1
    file_name = os.path.basename(file).removesuffix(".txt")
    rounded_score = round(yes_count / (yes_count + no_count) * 100, 2)
    print(f"{file_name},{rounded_score}")

#
# REBUILD
#

elif sys.argv[1] == "--rebuild":
  with open(sys.argv[2], "r", encoding = "utf-8") as yes:
    yes_paras = yes.readlines()
  with open(sys.argv[3], "r", encoding = "utf-8") as no:
    no_paras = no.readlines()

  yes_lines_tmp = []
  no_lines_tmp  = []
  yes_lines     = []
  no_lines      = []

  for para in yes_paras:
    yes_lines_tmp += split_paras(para)
  for para in no_paras:
    no_lines_tmp  += split_paras(para)

  for line in yes_lines_tmp:
    line = clean(line)
    if line != "":
      yes_lines.append(line)
  for line in no_lines_tmp:
    line = clean(line)
    if line != "":
      no_lines.append(line)

  big_training_data = []
  for line in no_lines:
    big_training_data.append([0, line])
  for line in yes_lines:
    big_training_data.append([1, line])
  random.seed(42)
  random.shuffle(big_training_data)


  # Not wanting to use a CSV, I asked Claude ...
  # it suggested that a pd.DataFrame is an
  # acceptable alternative to a CSV.

  df = pd.DataFrame(big_training_data, columns = ['label', 'text'])

  vectorizer = TfidfVectorizer()

  X = vectorizer.fit_transform(df['text'])
  y = df['label']

  X_train, X_test, y_train, y_test = train_test_split(
    X,
    y,
    test_size = 0.15,
    random_state = 42
  )

  model = LogisticRegression(
    random_state = 42,
    max_iter     = 1000, #1000,
    solver       = 'saga'
    # The docs suggest that 'saga' is faster.
  )

  model.fit(X_train, y_train)

  pickle.dump(vectorizer, open(vectorizer_file, "wb"))
  pickle.dump(model, open(model_file, "wb"))
