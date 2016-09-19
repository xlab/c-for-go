var t = $("#demo"),
    e = t.find(".terminal .cli"),
    t2 = e.find(".t2"),
    t3 = e.find(".t3"),
    t4 = e.find(".t4"),
    t5 = e.find(".t5"),
    t6 = e.find(".t6"),
    t7 = e.find(".t7"),
    t8 = e.find(".t8"),
    t9 = e.find(".t9"),
    t10 = e.find(".t10"),
    t11 = e.find(".t11"),
    d = t2.text(),
    d2 = t5.text(),
    d3 = t9.text()

function hide() {
    t2.hide(), t3.hide(), t4.hide(), t5.hide(),
        t6.hide(), t7.hide(), t8.hide(), t9.hide(),
        t10.hide(), t11.hide();
};

var t = function(t, e, n, r) {
    void 0 === n && (n = 50), t.delay(e).css({
        opacity: 0
    }).animate({
        opacity: 1
    }, {
        duration: n,
        start: function() {
            t.show()
        },
        complete: r
    })
};

function animate() {
    t(t2, 500),
        $({
            textLen: 0
        }).delay(500).animate({
            textLen: d.length
        }, {
            duration: 1e3,
            easing: "linear",
            start: function() {
                t2.text("").show()
            },
            step: function(t) {
                t2.text(d.slice(0, Math.floor(t)))
            }
        }),
        t(t3, 1500, 300),
        t(t4, 1800),
        t(t5, 1800),
        $({
            textLen: 0
        }).delay(1800).animate({
            textLen: d2.length
        }, {
            duration: 1e3,
            easing: "linear",
            start: function() {
                t5.text("").show()
            },
            step: function(t) {
                t5.text(d2.slice(0, Math.floor(t)))
            }
        }),
        t(t6, 2800),
        t(t7, 3600, 300),
        t(t8, 3900),
        t(t9, 3900),
        $({
            textLen: 0
        }).delay(3900).animate({
            textLen: d3.length
        }, {
            duration: 1e3,
            easing: "linear",
            start: function() {
                t9.text("").show()
            },
            step: function(t) {
                t9.text(d3.slice(0, Math.floor(t)))
            }
        }),
        t(t10, 4900, 300),
        t(t11, 4900)
}
