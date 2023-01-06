"use strict";
exports.__esModule = true;
exports.MultiPager2 = exports.MultiPager = void 0;
function printm(m) {
    console.log(m);
    console.log(typeof m);
    return 0;
}
printm("abc");
var MultiPager = /** @class */ (function () {
    function MultiPager() {
        this.fetch = function (opts) {
            console.log(opts);
        };
    }
    return MultiPager;
}());
exports.MultiPager = MultiPager;
new MultiPager().fetch({ more: true });
new MultiPager().fetch({ more: true, params: 1 });
var MultiPager2 = /** @class */ (function () {
    function MultiPager2() {
        this.fetch = function (opts) {
            console.log(opts);
        };
    }
    return MultiPager2;
}());
exports.MultiPager2 = MultiPager2;
new MultiPager2().fetch({ more: true });
new MultiPager2().fetch({ more: true, params: void (1) });
new MultiPager2().fetch({ more: true });
new MultiPager2().fetch({ more: true, params: 1 });
