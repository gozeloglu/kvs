/*
Package kvs is the in-memory key-value storage for Go. It has 2 basic features
which are Get and Set. Basically, Set adds key-value pair in map and Get returns
the value of the given key. If there is no such a key in map, it returns empty
string.

Create Kvs object once before calling Get and Set methods:

	db, _ := kvs.Create("", "users", 1*time.Minute)
	db.Set("foo", "bar")
	val := db.Get("foo")

Before diving into the package, you need to know that kvs package can be used
in 2 different ways. As a first option, it can be used as a package that can
be imported at the beginning of the code. Then, you can call the functions as
above. As another way, you can create an HTTP server and call endpoints as
defined in server.go. You can send requests to server as regular API call.

You can call endpoints like this:

	http://localhost:1234/set	--> POST
	Body:
	{
    	"data": [
        	{
            	"key": "joe",
            	"value": "13"
        	}
    	]
	}

	http://localhost:1234/get/joe	--> GET

	http://localhost:1234/save	--> PUT

With /set endpoint, you can add key-value pair(s) in memory, map. Multiple
key-value pairs are allowed in "data" array.

With /get/<key> endpoint, you can get the value of <key>. It returns a JSON
object that stores key, value, and result.

With /save endpoint, you can save the data that stores in memory to the disk.
It returns result message.

Default port is assigned as :1234 for this server. Actually, it is not used for
package usage, only is used for server usage.

While creating Kvs object, you need to specify database name. It is obligatory
to pass it to Create function. If it is not specified, Create function returns
error.

Here is the sample usage:

	// Db name is "users"
	db, err := kvs.Create(":1234", "users", 1*time.Minute)

	// Returns error
	db, err := kvs.Create(":1234", "", 1*time.Minute)

For the server usage, after creating Kvs object you need to call Open method to
start server. Here is the usage:

	func main() {
		db, err := kvs.Create(":1234", "users", 1*time.Minute)
		if err != nil {
			log.Fatalf(err.Error())
		}

		// Starts server on localhost:1234
		db.Open()
	}

In Create function, there is third parameter named as duration. It provides
saving data periodically in given interval time. For example, if you specify
2*time.Minute as parameter, the data in map would be saved every 2 minutes.
*/
package kvs
