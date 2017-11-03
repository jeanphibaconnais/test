var express = require('express');
var cors = require('cors');
const notifier = require('node-notifier');
var app = express();
const WindowsToaster = require('node-notifier').WindowsToaster;

var notifierWindows = new WindowsToaster({
    withFallback: false, // Fallback to Growl or Balloons?
    customPath: void 0 // Relative/Absolute path if you want to use your fork of SnoreToast.exe
});

app.use(cors());

app.get('/sendNotification', function (req, res) {

    // Object
    notifier.notify({
        'title': 'My notification',
        'message': 'Hello, there!'
    });

    res.send("Notification envoyée");

});

app.get('/sendNotificationMax', function (req, res) {
    var retour = "";

    notifier.notify({
        title: 'Notification',
        message: 'Voici une belle notification.',
        //icon: path.join(__dirname, 'coulson.jpg'), // Absolute path (doesn't work on balloons)
        sound: true, // Only Notification Center or Windows Toasters
        wait: true // Wait with callback, until user action is taken against notification
    }, function (err, response) {
        // Response is response from notification
    });

    notifier.on('click', function (notifierObject, options) {
        retour += "click";
    });

    notifier.on('timeout', function (notifierObject, options) {
        // Triggers if `wait: true` and notification closes
    });

    res.send(retour);

});


app.listen(8080);

console.log("The server is running on the 8080 port.");