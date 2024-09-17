/*
- Currently your company emits an HTTP redirect status code to send customers from an old webpage to a new one
- Now you want the ability to check how many customers are still trying to access the old page
- HTTP server log files are being saved in the logs diectory and some are compressed
- You want to see how many HTTP redirect codes you have been throwing
- ***Solution: write a function that gets an io.Reader and returns the total number of lines and the number of lines that have an
HTTP redirect directive
*/

package redirects