from jsonschema import validate

schema ={
		  "colors": [
		    {
		      "color": "black",
		      "category": "hue",
		      "type": "primary",
		      "code": {
		        "rgba": [255,255,255,1],
		        "hex": "#000"
		      }
		    },


		    {
		      "color": "white",
		      "category": "value",
		      "code": {
		        "rgba": [0,0,0,1],
		        "hex": "#FFF"
		      }
		    },
		    {
		      "color": "red",
		      "category": "hue",
		      "type": "primary",
		      "code": {
		        "rgba": [255,0,0,1],
		        "hex": "#FF0"
		      }
		    }
		  ]
		}
        # If no exception is raised by validate(), the instance is valid.
validate({"colors[0]" : "black","type": "primary", "code" : [0,0,1,0]}, schema)
# ValidationError: [0,0,1,0] is not a "code"