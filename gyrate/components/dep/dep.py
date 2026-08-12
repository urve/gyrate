#!/usr/bin/env python3

import warnings
warnings.filterwarnings("ignore", module="urllib3")

import spacy
import sys
import re
import os
import tqdm
import json
import sqlite3
from pathlib import Path

#nlp = spacy.load("en_core_web_sm")
#nlp = spacy.load("en_core_web_md")
nlp = spacy.load("en_core_web_lg")
# Should consider whether to make this a larger
# model to more accurately get the items ...

def file_to_sentences(file):
  final_sentences = []
  lines = file.readlines()
  for line in lines:
    sentences = re.split(r"[?!.]", line)
    for sentence in sentences:
      sentence = sentence.strip()
      if sentence == "" or sentence == "\"" or sentence == "'":
        continue
      else:
        final_sentences.append(sentence.lower())
  return final_sentences

def get_year(connection, table, file_name):
  file_year = connection.execute(f"SELECT year FROM {table} WHERE original_id = \"{file_name}\";").fetchone()
  return file_year

def argument_to_terms(argument):
  arguments = []
  if argument.is_file():
    with open(argument, 'r') as a:
      lines = file_to_sentences(a)
      for line in lines:
        nlp_line = nlp(line)[0].lemma_.lower()
        arguments.append(nlp_line)
  else:
    nlp_argument = nlp(str(argument))[0].lemma_.lower()
    arguments = [nlp_argument]
  return arguments

#nlp = spacy.load("en_core_web_sm")
ruler = nlp.get_pipe("attribute_ruler")
ruler.add(
  patterns = [[{"LOWER": "semen"}]],
  attrs = {"LEMMA": "semen"}
)

script_dir = os.path.dirname(os.path.abspath(__file__))
text_dir = Path(sys.argv[1])
db = Path(sys.argv[2])
table = sys.argv[3]
connection = sqlite3.connect(db)
terms_dir = Path(script_dir)/  "terms"
terms_dir.mkdir(parents=True, exist_ok=True)

argument = Path(sys.argv[4])
nlp_terms = []
terms = argument_to_terms(argument)
for term in terms:
  nlp_terms.append(term)

# Weird plurals in my current corpus
# that one may want to have dependency
# parsing for...
exceptions = {
    "tooth": "teeth",
    "foot": "feet",
    "abdomen": "abdoman",
    "think": "thought"
    }

for term in tqdm.tqdm(nlp_terms, desc="terms", position=0):
  termed_dir = terms_dir / term
  termed_dir.mkdir(parents=True, exist_ok=True)
  big_file = os.path.join(termed_dir, term + ".json")
  if os.path.isfile(big_file):
    continue
  year_dict = {}
  # year_dict has years as keys
  global_verbs = {}
  years = []
  for file in text_dir.iterdir():
    if file.name == ".DS_Store":
      continue
    file_name = os.path.basename(file).removesuffix(".txt")
    year = get_year(connection, table, file_name)[0]
    if year not in years:
      years.append(year)
    year_dict[year] = year_dict.get(year, {})

  for file in tqdm.tqdm(sorted(text_dir.iterdir()), desc=term, position=1, leave=False):
    if file.name == ".DS_Store":
      continue
    with open(file, 'r') as f:
      file_name = os.path.basename(file).removesuffix(".txt")
      year = get_year(connection, table, file_name)[0]
      sentences = file_to_sentences(f)

      for sentence in sentences:
        # The original handling of these variants was problematic;
        # Claude suggested a fix using this `if variant` logic.
        variant = exceptions.get(term)
        if variant:
          if not (re.search(term, sentence) or re.search(variant, sentence)):
            continue
        elif not re.search(term, sentence):
            continue
        doc = nlp(sentence)
        for token in doc:
          if token.lemma_.lower() == term and token.pos_ == "NOUN":
            # Originally `token.text.lower()`, but Claude
            # suggested lemmatization at this stage
            # to capture more items
            head = token.head
            if head.pos_ == "VERB":
              lem = head.lemma_
              # initialize or add to values for lem
              global_verbs[lem] = global_verbs.get(lem, 0) + 1
              year_dict[year][lem] = year_dict[year].get(lem, 0) + 1
  with open(big_file, 'w') as f:
    # write to .json file for both
    json.dump(
        dict(sorted(global_verbs.items(), key=lambda item: item[1], reverse=True)),
        f,
        indent=2
        )
  for year in years:
    year_file = os.path.join(termed_dir, term + "-" + str(year) + ".json")
    with open(year_file, 'w') as f:
      json.dump(
          dict(sorted(year_dict[year].items(), key=lambda item: item[1], reverse=True)),
          f,
          indent=2
          )

