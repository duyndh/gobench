(this["webpackJsonpgobench-ui"]=this["webpackJsonpgobench-ui"]||[]).push([[19],{622:function(e,t,a){"use strict";a.r(t);var n=a(126),c=a(0),i=a.n(c),l=a(38),r=a(27),m=a(45),s=a(59),o=Object(c.lazy)((function(){return a.e(20).then(a.bind(null,606))}));t.default=Object(r.i)(Object(l.c)((function(e){var t=e.application,a=e.dispatch;return{detail:t.detail,graphMetrics:t.graphMetrics,dispatch:a}}))((function(e){var t=e.detail,a=e.graph,l=e.graphMetrics,r=e.timestamp,u=e.dispatch,p=Object(c.useState)(s.d["1h"]),d=Object(n.a)(p,2),h=d[0],g=d[1],E=Object(m.get)(t,"status",""),b=!["finished","cancel"].includes(E);return Object(c.useEffect)((function(){a&&l.every((function(e){return e.graphId!==a.id}))&&u({type:"application/GRAPH_METRICS",payload:{id:a.id}})}),[a.id]),i.a.createElement(i.a.Fragment,null,i.a.createElement("div",{className:"graph"},i.a.createElement("div",{className:"graph-header"},i.a.createElement("h5",{title:a.id||"",className:"graph-title"},Object(m.get)(a,"title","")," (",Object(m.get)(a,"unit",""),")"),b&&i.a.createElement("div",{className:"options-group"},i.a.createElement("ul",{className:"time-range-options-list"},i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:h===s.d["5m"]?"active":"",onClick:function(){return g(s.d["5m"])}},"5m")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:h===s.d["15m"]?"active":"",onClick:function(){return g(s.d["15m"])}},"15m")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:h===s.d["30m"]?"active":"",onClick:function(){return g(s.d["30m"])}},"30m")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:h===s.d["1h"]?"active":"",onClick:function(){return g(s.d["1h"])}},"1h")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:h===s.d["12h"]?"active":"",onClick:function(){return g(s.d["12h"])}},"12h")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:h===s.d["24h"]?"active":"",onClick:function(){return g(s.d["24h"])}},"24h"))))),i.a.createElement(c.Suspense,{fallback:i.a.createElement("p",null,"Loading graph...")},i.a.createElement(o,{timeRange:h,graph:a,timestamp:r,unit:Object(m.get)(a,"unit","")}))))})))}}]);
//# sourceMappingURL=19.505ce039.chunk.js.map