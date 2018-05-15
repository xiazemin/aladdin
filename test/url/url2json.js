/**
 * Created by didi on 18/5/13.
 */
var urlToJson = function(){
    var ret = {};
    window.location.search.substr(1).replace(/(\w+)=(\w+)/ig, function(a, b, c){ret[b] = unescape(c);});
    return ret;
};
