#Building upon the code from the previous exercise

We are now going to get "I see you connected" to be written.
When we used bufio.NewScanner(), our code was reading from a io.Reader that never ended. 

We will now break out of the reading. 

Package bufio has the Scanner type. The Scanner type "provides a convinient interface for reading data". When you have a scanner type, you can call the SCAN method on it. Successive calls to the Scan method will step through the tokens (piece of data). The default token is a line. The Scanner type also has a TEXT method. When you call this method, you will be given the text from the current token. Here is how you will use it:

var scanner *bufio.Scanner = bufio.NewScanner(conn)

for scanner.Scan(){
    var ln string = scanner.Text()
    fmt.Println(ln)
}

Use this code to READ from an incoming connection and print the incoming text to standard out ( the terminal ).

When you "ln" line of text is equal to an empty string break out of the loop.
Run your code and goto localhost:8080 in **your browser. **

What do you find?