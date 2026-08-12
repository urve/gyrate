#!/usr/bin/env python3

import warnings
warnings.filterwarnings("ignore", module="urllib3")
# Claude suggested this importation because, at least
# in my environment, urllib3 tends to throw an error
# at runtime, but it appears to be an immaterial one.

from sentence_transformers import SentenceTransformer
import sys
import os
import re
from pathlib import Path
import pickle

def file_to_sentences(file):
  final_sentences = []
  lines = file.readlines()
  for line in lines:
    sentences = re.split(r'[?!.]', line)
    for sentence in sentences:
      sentence = sentence.strip()
      # skip blank sentences
      if sentence == "" or sentence == "\"" or sentence == "'":
        continue
      else:
        final_sentences.append(sentence)
  return final_sentences

script_dir = os.path.dirname(os.path.abspath(__file__))
pickles    = os.path.join(script_dir, "./pickles")

directory = Path(sys.argv[2])
query = " ".join(sys.argv[3:])

#transformer = "BAAI/bge-large-en-v1.5"
transformer = "BAAI/bge-small-en-v1.5"

transformer_sanitized = transformer.replace("/", "-")

model = SentenceTransformer(transformer)

embedded_query = model.encode(query)

for file in sorted(directory.iterdir()):
  if file.name == ".DS_Store":
    continue
  with open(file, "r") as f:
    file_name = os.path.basename(file).removesuffix(".txt")
    sentences = file_to_sentences(f)
    length = len(sentences)
    yes_score = 0
    # file name of the pickle
    pickled_embedding = os.path.join(pickles, f"{transformer_sanitized}_{file_name}.pickle")
    # rebuild
    if sys.argv[1] == "-r":
      os.makedirs(pickles, exist_ok=True)
      embeddings = model.encode(sentences)
      with open(pickled_embedding, "wb") as pickled:
        pickle.dump(embeddings, pickled, -1)
    # load
    elif sys.argv[1] == "-l":
      with open(pickled_embedding, "rb") as pickled:
        embeddings = pickle.load(pickled)
    score_list = model.similarity(embedded_query, embeddings)[0]
    for index, score in enumerate(score_list):
      # relatively arbitrary matching value of 0.6
      if score > 0.6:
        yes_score += 1
    print(f"{file_name},{(100 * yes_score / length):.2f}")

