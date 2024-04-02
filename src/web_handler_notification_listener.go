// THIS IS AN AUTOGENERATED FILE... See gen_web_handler.sh for details.
package qmq

import (
    "net/http"
    "fmt"
)

func Register_web_handler_notification_listener() {

    http.HandleFunc("js/qmq/notification_listener.js", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, `class NotificationListener {
    constructor() {}
    listenTo(topic, manager) {
        manager.addListener(topic, this);
        return this;
    }
    onNotification(topic, data, context) {}
};`)
    })
}
