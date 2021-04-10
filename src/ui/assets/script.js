/* mdwiki scipts */
let ui = {};
ui.api = {};
ui.dom = {};
ui.browser = {};
ui.editor = {};

ui.api = (function () {
    function handleResponse(response) {
        return response.text().then(text => {
            const data = text && JSON.parse(text);

            if (!response.ok) {
                const error = (data && data.message) || response.statusText;
                return Promise.reject(error);
            }

            return data;
        });
    }

    return {
        post: function (url, data) {
            return fetch(url, {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                redirect: "follow",
                body: JSON.stringify(data)
            }).then(handleResponse);
        }
    }
})();

ui.dom = (function () {
    return {
        el: function (query) {
            return document.querySelector(query)
        },
        content: function (query, content) {
            document.querySelector(query).innerHTML = content
        },
        onLoad: function (listener) {
            window.addEventListener("DOMContentLoaded", listener)
        },
        onClick: function (query, listener) {
            document.querySelector(query).addEventListener("click", listener)
        },
        onKey: function (query, listener) {
            document.querySelector(query).addEventListener("keydown", listener)
        },
        onPaste:  function (query, listener) {
            document.querySelector(query).addEventListener("paste", listener)
        },
        hide: function (query) {
            document.querySelector(query).style.display = "none"
        },
        show: function (query) {
            document.querySelector(query).style.display = "block"
        },
        toggle: function (query) {
            let e = document.querySelector(query)
            window.getComputedStyle(e).display === "block" ?
                this.hide(e) : this.show(e);
        }
    }
})();

ui.browser = (function () {
    return {
        redirect: function (location) {
            window.location.replace(location)
        }
    }
})();

ui.editor  = function (query) {
    ui.dom.onKey(query, function (e) {
        if (e.keyCode == 9) {
            e.preventDefault()
            document.execCommand('insertHTML', false, '&#009');
        }
        if (e.keyCode == 13) {
            e.preventDefault()
            document.execCommand('insertLineBreak')
        }
    })
    ui.dom.onPaste(query, function (e) {
        e.preventDefault();
        const text = e.clipboardData.getData("text/plain");
        document.execCommand("insertText", false, text);
    })

    const editor = ui.dom.el(query)
    editor.contentEditable = true
    editor.spellcheck = false

    return {
        focus: function() {
            editor.focus()
        },
        getText: function() {
            return editor.innerText
        }
    }
};


