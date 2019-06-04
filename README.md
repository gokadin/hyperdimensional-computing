# Hyperdimensional Computing
The goal is to demonstrate the concept of using hyperdimensional vectors with a simple example. 

## What is it?
When working in hyperspace (say 10000 dimensions), some interesting properties arise. Example:
- We can pick any number of random vector from that space and each one will be approximately orthogonal to all other ones previously picked. 
- We can ADD two vectors together to obtain a third vector that is <b>similar</b> to the original ones. 
- We can multiply two vectors to obtain a third vector that is <b>dissimilar</b> to the original ones. 
- We can reconstruct the meaning of a vector by taking any 60% of its values (will simply be noisier).

Using these properties, we can encode patterns, correlate them and compare them easily, while distributing information, therefore having a highly redundant and robust system. 

## Example
The  example will show how the algorithm can guess the language of any phrase or text in the set of languages that it has previously learned. 

The algorithm for this example is as follows:
1. Encode  each letter: generate a random 10k vector for each letter and store them in a map. 
2. Encode trigrams: trigrams are sequences of 3 letters. Example, "something" will generate trigrams: "som", "ome", "met", "eth", "thi", "hin", "ing".
3. Encode language profile: add each of the trigrams to form a <b>similar</b> vector to all of the trigram vectors. 
4. Repeat for any other language and for the test string. 
5. Compare the test vector with each language vector using cosine similarity. 

# Resources
- Paper on hyperdimensional computing by Pentti Kanerva: http://www.rctn.org/vs265/kanerva09-hyperdimensional.pdf
