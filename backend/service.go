package backend

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"
)

type Repository struct{}

const USER_DETAILS_URL = "http://uinames.com/api/"
const stringUrl = "http://api.icndb.com/jokes/random?limitTo=[nerdy]"

//TODO : better error response
func (r Repository) GetJoke() string {
	response_message := " Why newton why ?"
	// no harm calling const url
	response, err := http.Get(USER_DETAILS_URL)
	if err != nil {
		response_message = err.Error()
		return response_message
	}
	data, err := ioutil.ReadAll(response.Body)
	if err!=nil{
		response_message = err.Error()
		return response_message
	}
	var _name Person
	err = json.Unmarshal(data, &_name)
	if err!=nil{
		response_message = err.Error()
		return response_message
	}
	//build url from response
	params := url.Values{}
	params.Add("firstName", _name.Name)
	params.Add("lastName", _name.Surname)
	//TODO : make this slice as per url it must be slice
	params.Add("limitTo", "[nerdy]")
	baseUrl, err := url.Parse(stringUrl)
	if err!=nil{
		response_message = err.Error()
		return response_message
	}
	baseUrl.RawQuery = params.Encode()
	jokeResponse, err := http.Get(baseUrl.String())
	if err != nil {
		response_message = err.Error()
		return response_message
	} else {
		joke, err := ioutil.ReadAll(jokeResponse.Body)
		if err!=nil{
			response_message = err.Error()
			return response_message
		}
		var _joke JokeObj
		err = json.Unmarshal(joke, &_joke)
		if err!=nil{
			response_message = err.Error()
			return response_message
		}
		response_message = _joke.Value.Joke
	}


	return response_message
}
