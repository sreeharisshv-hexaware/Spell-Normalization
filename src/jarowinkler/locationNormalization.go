package main

/*
    Location Normalization
    ==================================

    Read a list of places and spell correct them using string distance metrics. 
*/

import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  "github.com/xrash/smetrics"
  "sort"
)

type localityObject struct {
  Chennai []string 
}

var chennai_locality_map map[string] float64;

func locationNormalization(misspelled_word  string) {
  /*
      @param: misspelled_word   
      Function to compute and return corrected word based on distance-measure
  */
  filePath := "./chennai_locality_list.json";
  file, err1 := ioutil.ReadFile( filePath )

  if err1 != nil {
      fmt.Printf( "// error while reading file %s\n", filePath )
      fmt.Printf("File error: %v\n", err1)
      os.Exit(1)
  }

  var chennai_locality localityObject

  err2 := json.Unmarshal(file, &chennai_locality)
  if err2 != nil {
    fmt.Println("error:", err2)
    os.Exit(1)
  }
 
  chennai_locality_map = make(map[string]float64)

  for _,i:= range chennai_locality.Chennai{
    chennai_locality_map[i]=smetrics.JaroWinkler(i, misspelled_word , 0.7, 4);
  }  

  var data []string = sortScoreValue(chennai_locality_map, 1)
  
  for x := range data{
    fmt.Println(data[x])
  }
}

func sortScoreValue(scored_map map[string]float64, number_similar_locality int) []string {

  /*
      @param: scored_map - Map containing correct locality names and scores returned by Jaro-Winkler Algorithm
      @param: number_similar_locality - number of similar locality names to return 
      Function to reverse sort the word scores and return the top n matches
  */
  temp_score_map := map[float64][]string{}
  var locality_score []float64

  for locality, score := range scored_map {
          temp_score_map[score] = append(temp_score_map[score], locality)
  }
  
  for loc_score := range temp_score_map {
          locality_score = append(locality_score, loc_score)
  }
  
  sort.Sort(sort.Reverse(sort.Float64Slice(locality_score)))
  matched_locality_list := make([]string, number_similar_locality)
  
  for i:=0; i< number_similar_locality; i++ {
    matched_locality_list[i]=temp_score_map[locality_score[i]][0]
  }

  return matched_locality_list
}