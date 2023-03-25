// Import required packages and rename one of them to avoid naming conflict

package main

import (
	"flag"
	"fmt"
	"example.com/m/v2/conv"
	myfunfacts "example.com/m/v2/funfacts" // Rename the imported package to avoid naming conflict
	"regexp"
	"strconv"

)

// Declare global variables to hold the values passed through command line arguments
var (
	fahr      float64
	celsius   float64
	kelvin    float64
	out       string
	funfacts  string
	tempSkala string
)

// init function to set up flags for command line arguments
func init() {
	flag.Float64Var(&celsius, "celsius", 0.0, "Temperature in Celsius")
	flag.Float64Var(&fahr, "fahr", 0.0, "Temperature in Fahrenheit")
	flag.Float64Var(&kelvin, "kelvin", 0.0, "Temperature in Kelvin")
	flag.StringVar(&out, "out", "celsius", "Output temperature scale (celsius, fahrenheit, kelvin)")
	flag.StringVar(&funfacts, "funfacts", "", "Get fun facts about sun, luna, or terra") // Add a flag to get fun facts
}

// main function to handle the logic
func main() {
// Parse the command line arguments
	flag.Parse()
// If funfacts flag is not present, convert the temperature based on output temperature scale
	if funfacts != "" {
    	funFacts := myfunfacts.GetFunFacts(funfacts)
    	for i, fact := range funFacts {
    		fmt.Println(i+1, fact)
    	}
    } else {
    	var convValue float64
    	var convFunc func(float64) float64
    	switch out {
    	case "celsius":
    		convFunc = conv.FahrenheitToCelsius
    		convValue = fahr
    		if kelvin != 0 {
    			convFunc = conv.KelvinToCelsius
    			convValue = kelvin
    		}
    	case "fahrenheit":
    		convFunc = conv.CelsiusToFahrenheit
    		convValue = celsius
    		if kelvin != 0 {
    			convFunc = conv.FahrenheitToKelvin
    			convValue = fahr
    		}
    	case "kelvin":
    		convFunc = conv.CelsiusToKelvin
    		convValue = celsius
    		if fahr != 0 {
    			convFunc = conv.FahrenheitToKelvin
    			convValue = fahr
    		}
    		// If the output temperature scale is not valid, print an error message and exit
    	default:
    		fmt.Println("Invalid output temperature scale:", out)
    		return
    	}

        // Format the result string based on the result value
    	result := convFunc(convValue)
    	var resultStr string
    	if result < 1000000 {
    		if result == float64(int(result)) {
    			resultStr = strconv.Itoa(int(result))
    		} else {
    			resultStr = fmt.Sprintf("%.2f", result)
    		}
    	} else {
    		resultStr = strconv.FormatFloat(result, 'f', 0, 64)
             re := regexp.MustCompile(`(\d)(\d{3})($|\.)`)
             resultStr = re.ReplaceAllString(resultStr, "$1 $2$3")

    	}
    	// Print the result string
    	fmt.Println(resultStr)
    }
    }