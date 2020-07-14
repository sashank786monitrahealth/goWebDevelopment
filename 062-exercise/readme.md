<h2>Building upon the code from the previous problem:</h2>

<p>extract the code you wrote to READ from the connection using bufio.NewScanner into its own 
    function called <b>serve</b>
</p>

<p>Pass the connection of type net.Conn as an argument into this function. </p>

<p>Add <b>go</b> in front of the call to <b>serve</b> to enable concurrency and multiple connections.</p>