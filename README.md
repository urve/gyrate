# gyrate

## Get started
```bash
git clone https://github.com/urve/gyrate
cd gyrate/gyrate
bash SETUP
./gyrate --help
```

See [funding and license information](#Funding-and-license-information) how this project came to be.

## Usage

### Overview

This is a wrapper for a suite of text analysis tools for digital humanists attempting to classify and compare texts in a large corpus. Three different classification methods -- naive Bayesian classification, logistic regression, and a dirty vector embedding -- are provided.

The data can be difficult to interpret, so a few bridging utilities are included. These include connecting to a sqlite3 database to get years, normalize CSVs, parse dependencies (associations of verbs to a noun), and detect duplicates.

### Split-and-classify

`sc CANNIBALIZED_FILE FIRST_CLASS.txt SECOND_CLASS.txt`

`sc` takes a `CANNIBALIZED_FILE` and iterates over its sentences. You can press `y` or `n` for each sentence for whether it classifies as `FIRST_CLASS` (`y`) or `SECOND_CLASS`.

This returns two `.txt` files in the `classifications/` directory: `FIRST_CLASS.txt` and `SECOND_CLASS.txt`. These can then be fed into the Bayesian or logistic regression classifiers. For ease of use, and since users may be classifying several different texts over time, each `CANNIBALIZED_FILE`'s classified sentences are preceded by a line containing many different `─` symbols and an identifier for the file.

### Duplicate detection

`./gyrate --duplicates DIRECTORY`

This determines whether each file in the DIRECTORY contains a duplicate. Duplicates are found using TF-IDF cosine similarity of documents; if two files contain a similarity above a certain treshold (currently `0.95`), then it returns a `.csv` to STDOUT of the format `ID_1,ID_2,similarity`. One can then exclude these items for text analysis and classification.

(Note: A switch on each classifier excluding duplicate files, according to whether a file is an earlier or a later edition of the same text, is in the works.)

### Dependency parsing

`./gyrate --dependency DIRECTORY DB TABLE TERM`

Assuming that a sqlite3 database `DB` contains a table `TABLE`, which includes both `original_id` (corresponding to filenames) and `year` column, and this returns a series of `.json` files in `components/dep/terms`. These files are created based on the `TERM`, which is either a single noun passed directly into the command or a file of newline-delimited single nouns. Because this parsing can, at times, take quite a long time if the `TERM` is common, or if the `DIRECTORY` is large, progress bars are shown.

A directory `components/dep/terms/TERM` is created. For each `year` that are represented in the `DIRECTORY`, a file of the form `TERM-year.json` is created, which is of the form `"verb": count`. Each `verb` is a verb associated with the noun `TERM`, either in active or passive voice. A file `TERM.json` is also created which represents appearances of `TERM` across the corpus.

This currently uses `spacy`'s `en_core_web_lg`, which is large. You can adjust the
model accordingly.

### Year bridge

`./gyrate --year-bridge DB TABLE`

This can only be used in pipes. Assuming that a sqlite3 database `DB` contains a table, which includes both `original_id` and `year` columns, this returns a `.csv` of the form `ID,YEAR,SCORE`. This is useful in piping items around to utilities like `xan`, which can then, for example, return averages of scores, aggregated by year.

For instance, `./gyrate --bayes-all DIRECTORY YES_FILE NO_FILE | ./gyrate --year-bridge DB TABLE | xan sort -ns 1 | xan groupby -n '1' --along-cols '2' 'mean(_)'` returns Bayesian classification scores of a `DIRECTORY` ordered by `year`.

### Comparisons

`./components/values/difference CSV_1 CSV_2`

Assuming that `CSV_1` and `CSV_2` are of the form `ID,SCORE`, this returns a `.csv` that shows their difference (`SCORE_1 - SCORE_2`). This might be especially useful with the use of `--normalize` to see how different classification methods score texts. On my corpus, these are not terribly different.

### Normalize

`./gyrate --normalize CSV`

Sometimes, the actual scores of each classification method are not that important; what matters is their *relative* scores. In cases such as these, `--normalize` scales the items of the `CSV` so the lowest value `ID` is 0, and the highest value `ID` is 100. These are then returned in the `ID,SCALED_VALUE` format used elsewhere.

### Bayes

`./gyrate --bayes FILE YES_FILE NO_FILE`

Performs a naive Bayesian classification on the `FILE`. The number of sentences that are classified as YES, based on the `YES_FILE` and in proportion to the total number of classified sentences as either YES or NO, is returned. Specifically, a one-row `.csv` of the format `ID,SCORE` is returned to STDOUT.

To get the scores of an entire directory in a single `.csv`, `./gyrate --bayes-all DIRECTORY YES_FILE NO_FILE` can be used. Because this uses GNU `parallel`, the results are not guaranteed to be (and are usually not) in alphabetical or ID order.

### Logistic regression

`./gyrate --logistic-rebuild YES_FILE NO_FILE`

This returns two `.pickle` files -- `model.pickle` and `vectorizer.pickle` -- in `components/lr`, which are a logistic regression model and vectorizer, respectively, trained on the `YES_FILE` and the `NO_FILE`.

One can then use `./gyrate --logistic FILE`, which automatically pulls the model and vectorizer, to return logistic regression scores for the `FILE` in a `.csv` format of `ID,SCORE`.

Like Bayesian classification, the option to use GNU `parallel` to return a `.csv` (not necessarily in ID order) to STDOUT can be obtained with `./gyrate --logistic-all DIRECTORY`.

### Embedding

`./gyrate --embed-rebuild DIRECTORY QUERY`

Supposing that `QUERY` is a quotation-mark-enclosed string, the `--embed-rebuild` creates a series of `.pickles` in `components/embed/pickles` that correspond to a matrix of embeddings of each sentence in each file in `DIRECTORY`. Because there is no meaningful overhead in doing so, after rebuilding, `--embed-rebuild` also performs `--embed` on the `QUERY` (see below).

`./gyrate --embed DIRECTORY QUERY` returns to STDOUT a `.csv`, like `--bayes-all` or `--logistic-all` containing the proportion of sentences with a cosine similarity to `QUERY` above a certain threshold. That threshold is currently `0.6` but can be altered in `components/embed/embed.py`.

### Examples

A few barebones examples are given in `./gyrate --examples`. These show some ways that `gyrate` can be used in pipelines.

### WIP

* Make Bayes, logistic, and vector versions all be able to do sentence scoring
  * ~~Bayes~~
  * logistic
  * vector

* `--gazetteer`: a gazetteer to track places mentioned in works across the corpus, perhaps distinguishable by year

## Requirements

### Utilities
- GNU `parallel`
- `xan`

Installation of these depends on your operating system. On macOS, you can use `brew`.

### Go modules
- `github.com/jbrukh/bayesian`
- `github.com/mattn/go-sqlite3`
- `github.com/fatih/color`

### Python packages
- `pandas`
- `scikit_learn`
- `sentence_transformers`
- `spacy`
- `tqdm`

## Funding and license information

Creation of `gyrate` was supported by the [Center for Digital Research in the Humanities](https://cdrh.unl.edu/) at the [University of Nebraska–Lincoln](https://unl.edu/) through the 2026 Summer Fellowship. More information can be read in [Nebraska Today](https://news.unl.edu/article/summer-2026-digital-humanities-fellows-announced/).

This software is released under the GPLv3+: GPL 3.0 or later. See `COPYING` for details.
