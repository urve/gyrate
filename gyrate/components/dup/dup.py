#!/usr/bin/env python3

import sys
from pathlib import Path
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity

directory = Path(sys.argv[1])

corpus = []
files = []

for file in directory.iterdir():
  if file.name == ".DS_Store":
    continue
  with open(file, "r") as f:
    files.append(file)
    text_array = f.readlines()
    text = ""
    for line in text_array:
      line = line.replace("\n", " ")
      text += line
    corpus.append(text)

#print(len(corpus))

vectorizer = TfidfVectorizer(stop_words='english', analyzer='word', ngram_range=(1,2))

# Originally, I was using a nested loop which pairwise computed the cosine similarity
# between two documents. For just a few hundred books, that ended up taking several
# hours for the computations to finish.
# To speed this up, Claude suggested that a matrix be used instead of pairwise
# comparison of two vectors ... and it is much faster now.

matrix = vectorizer.fit_transform(corpus)
similarity_matrix = cosine_similarity(matrix)

for index in range(len(files)):
  for second_index in range(index + 1, len(files)):
  # index + 1 allows us to skip already-checked pairs
  # and prevents case where index = second_index
    if similarity_matrix[index][second_index] >= 0.95:
      print(f"{files[index].name},{files[second_index].name},{similarity_matrix[index][second_index]}")
