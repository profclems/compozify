(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[189],{2685:function(e,t,n){Promise.resolve().then(n.t.bind(n,7560,23)),Promise.resolve().then(n.t.bind(n,5502,23)),Promise.resolve().then(n.t.bind(n,9949,23)),Promise.resolve().then(n.t.bind(n,6966,23)),Promise.resolve().then(n.t.bind(n,9825,23)),Promise.resolve().then(n.t.bind(n,9915,23)),Promise.resolve().then(n.bind(n,7095)),Promise.resolve().then(n.bind(n,188)),Promise.resolve().then(n.t.bind(n,6513,23)),Promise.resolve().then(n.bind(n,783)),Promise.resolve().then(n.bind(n,459)),Promise.resolve().then(n.bind(n,8301)),Promise.resolve().then(n.t.bind(n,5552,23))},7463:function(e,t){"use strict";let n;Object.defineProperty(t,"__esModule",{value:!0}),function(e,t){for(var n in t)Object.defineProperty(e,n,{enumerable:!0,get:t[n]})}(t,{DOMAttributeNames:function(){return r},isEqualNode:function(){return a},default:function(){return o}});let r={acceptCharset:"accept-charset",className:"class",htmlFor:"for",httpEquiv:"http-equiv",noModule:"noModule"};function l(e){let{type:t,props:n}=e,l=document.createElement(t);for(let e in n){if(!n.hasOwnProperty(e)||"children"===e||"dangerouslySetInnerHTML"===e||void 0===n[e])continue;let a=r[e]||e.toLowerCase();"script"===t&&("async"===a||"defer"===a||"noModule"===a)?l[a]=!!n[e]:l.setAttribute(a,n[e])}let{children:a,dangerouslySetInnerHTML:o}=n;return o?l.innerHTML=o.__html||"":a&&(l.textContent="string"==typeof a?a:Array.isArray(a)?a.join(""):""),l}function a(e,t){if(e instanceof HTMLElement&&t instanceof HTMLElement){let n=t.getAttribute("nonce");if(n&&!e.getAttribute("nonce")){let r=t.cloneNode(!0);return r.setAttribute("nonce",""),r.nonce=n,n===e.nonce&&e.isEqualNode(r)}}return e.isEqualNode(t)}function o(){return{mountedInstances:new Set,updateHead:e=>{let t={};e.forEach(e=>{if("link"===e.type&&e.props["data-optimized-fonts"]){if(document.querySelector('style[data-href="'+e.props["data-href"]+'"]'))return;e.props.href=e.props["data-href"],e.props["data-href"]=void 0}let n=t[e.type]||[];n.push(e),t[e.type]=n});let r=t.title?t.title[0]:null,l="";if(r){let{children:e}=r.props;l="string"==typeof e?e:Array.isArray(e)?e.join(""):""}l!==document.title&&(document.title=l),["meta","base","link","style","script"].forEach(e=>{n(e,t[e]||[])})}}}n=(e,t)=>{let n=document.getElementsByTagName("head")[0],r=n.querySelector("meta[name=next-head-count]"),o=Number(r.content),s=[];for(let t=0,n=r.previousElementSibling;t<o;t++,n=(null==n?void 0:n.previousElementSibling)||null){var i;(null==n?void 0:null==(i=n.tagName)?void 0:i.toLowerCase())===e&&s.push(n)}let u=t.map(l).filter(e=>{for(let t=0,n=s.length;t<n;t++){let n=s[t];if(a(n,e))return s.splice(t,1),!1}return!0});s.forEach(e=>{var t;return null==(t=e.parentNode)?void 0:t.removeChild(e)}),u.forEach(e=>n.insertBefore(e,r)),r.content=(o-s.length+u.length).toString()},("function"==typeof t.default||"object"==typeof t.default&&null!==t.default)&&void 0===t.default.__esModule&&(Object.defineProperty(t.default,"__esModule",{value:!0}),Object.assign(t.default,t),e.exports=t.default)},6513:function(e,t,n){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),function(e,t){for(var n in t)Object.defineProperty(e,n,{enumerable:!0,get:t[n]})}(t,{handleClientScriptLoad:function(){return p},initScriptLoader:function(){return h},default:function(){return x}});let r=n(7023),l=n(5966),a=r._(n(6259)),o=l._(n(8471)),s=n(5314),i=n(7463),u=n(2815),d=new Map,c=new Set,f=["onLoad","onReady","dangerouslySetInnerHTML","children","onError","strategy"],m=e=>{let{src:t,id:n,onLoad:r=()=>{},onReady:l=null,dangerouslySetInnerHTML:a,children:o="",strategy:s="afterInteractive",onError:u}=e,m=n||t;if(m&&c.has(m))return;if(d.has(t)){c.add(m),d.get(t).then(r,u);return}let p=()=>{l&&l(),c.add(m)},h=document.createElement("script"),v=new Promise((e,t)=>{h.addEventListener("load",function(t){e(),r&&r.call(this,t),p()}),h.addEventListener("error",function(e){t(e)})}).catch(function(e){u&&u(e)});for(let[n,r]of(a?(h.innerHTML=a.__html||"",p()):o?(h.textContent="string"==typeof o?o:Array.isArray(o)?o.join(""):"",p()):t&&(h.src=t,d.set(t,v)),Object.entries(e))){if(void 0===r||f.includes(n))continue;let e=i.DOMAttributeNames[n]||n.toLowerCase();h.setAttribute(e,r)}"worker"===s&&h.setAttribute("type","text/partytown"),h.setAttribute("data-nscript",s),document.body.appendChild(h)};function p(e){let{strategy:t="afterInteractive"}=e;"lazyOnload"===t?window.addEventListener("load",()=>{(0,u.requestIdleCallback)(()=>m(e))}):m(e)}function h(e){e.forEach(p),function(){let e=[...document.querySelectorAll('[data-nscript="beforeInteractive"]'),...document.querySelectorAll('[data-nscript="beforePageRender"]')];e.forEach(e=>{let t=e.id||e.getAttribute("src");c.add(t)})}()}function v(e){let{id:t,src:n="",onLoad:r=()=>{},onReady:l=null,strategy:i="afterInteractive",onError:d,...f}=e,{updateScripts:p,scripts:h,getIsSsr:v,appDir:x,nonce:b}=(0,o.useContext)(s.HeadManagerContext),g=(0,o.useRef)(!1);(0,o.useEffect)(()=>{let e=t||n;g.current||(l&&e&&c.has(e)&&l(),g.current=!0)},[l,t,n]);let y=(0,o.useRef)(!1);if((0,o.useEffect)(()=>{!y.current&&("afterInteractive"===i?m(e):"lazyOnload"===i&&("complete"===document.readyState?(0,u.requestIdleCallback)(()=>m(e)):window.addEventListener("load",()=>{(0,u.requestIdleCallback)(()=>m(e))})),y.current=!0)},[e,i]),("beforeInteractive"===i||"worker"===i)&&(p?(h[i]=(h[i]||[]).concat([{id:t,src:n,onLoad:r,onReady:l,onError:d,...f}]),p(h)):v&&v()?c.add(t||n):v&&!v()&&m(e)),x){if("beforeInteractive"===i)return n?(a.default.preload(n,f.integrity?{as:"script",integrity:f.integrity}:{as:"script"}),o.default.createElement("script",{nonce:b,dangerouslySetInnerHTML:{__html:"(self.__next_s=self.__next_s||[]).push("+JSON.stringify([n])+")"}})):(f.dangerouslySetInnerHTML&&(f.children=f.dangerouslySetInnerHTML.__html,delete f.dangerouslySetInnerHTML),o.default.createElement("script",{nonce:b,dangerouslySetInnerHTML:{__html:"(self.__next_s=self.__next_s||[]).push("+JSON.stringify([0,{...f}])+")"}}));"afterInteractive"===i&&n&&a.default.preload(n,f.integrity?{as:"script",integrity:f.integrity}:{as:"script"})}return null}Object.defineProperty(v,"__nextScript",{value:!0});let x=v;("function"==typeof t.default||"object"==typeof t.default&&null!==t.default)&&void 0===t.default.__esModule&&(Object.defineProperty(t.default,"__esModule",{value:!0}),Object.assign(t.default,t),e.exports=t.default)},6966:function(e,t,n){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),function(e,t){for(var n in t)Object.defineProperty(e,n,{enumerable:!0,get:t[n]})}(t,{suspense:function(){return l},NoSSR:function(){return a}}),n(7023),n(8471);let r=n(6943);function l(){let e=Error(r.NEXT_DYNAMIC_NO_SSR_CODE);throw e.digest=r.NEXT_DYNAMIC_NO_SSR_CODE,e}function a(e){let{children:t}=e;return t}},188:function(e,t,n){"use strict";n.r(t),n.d(t,{default:function(){return i}});var r=n(5001),l=n(8471),a=n(7190);async function o(e,t){navigator.clipboard.writeText(e)}var s=n(7271);function i(e){let{value:t,className:n,src:i,...u}=e,[d,c]=(0,l.useState)(!1);return(0,l.useEffect)(()=>{setTimeout(()=>{c(!1)},2e3)},[d]),(0,r.jsxs)("button",{className:(0,a.cn)("relative z-20 inline-flex h-8 items-center justify-center rounded-md border-neutral-200 p-2 text-sm font-medium text-neutral-900 transition-all hover:bg-neutral-100 focus:outline-none dark:text-neutral-100 dark:hover:bg-neutral-800",n),onClick:()=>{o(t,{component:i}),c(!0)},...u,children:[(0,r.jsx)("span",{className:"sr-only",children:"Copy"}),d?(0,r.jsx)(s.dZ6,{className:"h-4 w-auto"}):(0,r.jsx)(s.Bp4,{className:"h-4 w-auto"})]})}},7095:function(e,t,n){"use strict";n.r(t);var r=n(5001),l=n(8471),a=n(9279),o=n.n(a),s=n(7190);let i=(0,l.forwardRef)((e,t)=>{let{className:n,children:l,...a}=e;return(0,r.jsxs)(o(),{ref:t,className:(0,s.cn)("group relative cursor-pointer px-1 py-0.5",n),...a,children:[l,(0,r.jsx)("span",{className:"absolute inset-x-0 bottom-0 block bg-zinc-900 h-[3px] w-0 rounded-md transition-width group-hover:w-full dark:bg-white"})]})});t.default=i},783:function(e,t,n){"use strict";n.r(t),n.d(t,{DocsPaginate:function(){return m},getPagerForDocs:function(){return p}});var r=n(5001),l=n(9279),a=n.n(l),o=n(8471),s=n(8567),i=n(9754);let u=(0,s.j)("active:scale-95 inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-neutral-400 focus:ring-offset-2 dark:hover:bg-neutral-800 dark:hover:text-neutral-100 disabled:opacity-50 dark:focus:ring-neutral-400 disabled:pointer-events-none dark:focus:ring-offset-neutral-900 data-[state=open]:bg-neutral-100 dark:data-[state=open]:bg-neutral-800",{variants:{variant:{default:"bg-neutral-900 text-white hover:bg-neutral-700 dark:bg-neutral-50 dark:text-neutral-900",destructive:"bg-red-500 text-white hover:bg-red-600 dark:hover:bg-red-600",outline:"bg-transparent border border-neutral-200 hover:bg-neutral-100 dark:border-neutral-700 dark:text-neutral-100",subtle:"bg-neutral-100 text-neutral-900 hover:bg-neutral-200 dark:bg-neutral-700 dark:text-neutral-100",ghost:"bg-transparent hover:bg-neutral-100 dark:hover:bg-neutral-800 dark:text-neutral-100 dark:hover:text-neutral-100 data-[state=open]:bg-transparent dark:data-[state=open]:bg-transparent",link:"bg-transparent dark:bg-transparent underline-offset-4 hover:underline text-neutral-900 dark:text-neutral-100 hover:bg-transparent dark:hover:bg-transparent"},size:{default:"h-10 py-2 px-4",sm:"h-9 px-2 rounded-md",lg:"h-11 px-8 rounded-md"}},defaultVariants:{variant:"default",size:"default"}}),d=o.forwardRef((e,t)=>{let{className:n,variant:l,size:a,...o}=e;return(0,r.jsx)("button",{className:(0,i.Z)(u({variant:l,size:a,className:n})),ref:t,...o})});d.displayName="Button";var c=n(7190),f=n(7271);function m(e){var t,n;let{docs:l,activeDocs:s}=e,i=p(l,s),d=function(e){let t=arguments.length>1&&void 0!==arguments[1]&&arguments[1],[n,r]=(0,o.useState)(t);return(0,o.useEffect)(()=>{let t=!0,n=window.matchMedia(e),l=()=>{t&&r(n.matches)};return n.addListener(l),r(n.matches),()=>{t=!1,n.removeListener(l)}},[e]),n}("(max-width: 425px)");return(0,r.jsxs)("div",{className:"relative flex flex-row items-center justify-between py-5",children:[i&&(null===(t=i.prev)||void 0===t?void 0:t.slug)&&(0,r.jsxs)(a(),{href:i.prev.slug,className:(0,c.cn)(u({variant:"outline"}),"absolute left-0"),children:[(0,r.jsx)(f.DEl,{className:"mr-2 h-4 w-4"}),(0,r.jsx)("span",{className:"overflow-hidden text-ellipsis whitespace-nowrap max-lg:max-w-[10rem] max-sm:text-xsm",children:d?"Previous":i.prev.title})]}),i&&(null===(n=i.next)||void 0===n?void 0:n.slug)&&(0,r.jsxs)(a(),{href:i.next.slug,className:(0,c.cn)(u({variant:"outline"}),"absolute right-0"),children:[(0,r.jsx)("span",{className:"overflow-hidden text-ellipsis whitespace-nowrap max-lg:max-w-[10rem] max-md:text-xsm",children:d?"Next":i.next.title}),(0,r.jsx)(f.MOd,{className:"ml-2 h-4 w-4"})]})]})}function p(e,t){let n=[null,...e,null],r=n.findIndex(e=>t.slug===(null==e?void 0:e.slug)),l=0!==r?n[r-1]:null,a=r!==n.length-1?n[r+1]:null;return{prev:l,next:a}}},459:function(e,t,n){"use strict";n.r(t),n.d(t,{DocsTableOfContents:function(){return m}});var r=n(5001),l=n(8471),a=n(9279),o=n.n(a),s=n(9756),i=n(9754),u=n(7271);function d(e){let{className:t,disableOnRoutes:n,disableOnLayouts:a}=e,o=(0,s.usePathname)(),d=function(){let[e,t]=(0,l.useState)(0);return(0,l.useEffect)(()=>{function e(){t(window.scrollY)}return window.addEventListener("scroll",e),()=>{window.removeEventListener("scroll",e)}},[]),e}();return(0,r.jsxs)("button",{className:(0,i.Z)(t,"group/link relative inline-flex items-center space-x-2 pb-1.5 pl-0.5 uppercase text-neutral-600 dark:text-neutral-400",n&&n.map(e=>e===o&&"hidden"),a&&a.map(e=>o.startsWith(e)&&"hidden"),d<100&&"hidden"),onClick:()=>{window.scrollTo({top:0,behavior:"smooth"})},children:[(0,r.jsx)(u.sIl,{className:"h-3 w-auto"}),(0,r.jsx)("span",{children:"Scroll to top"})]})}var c=n(7190),f=n(2162);function m(e){let{toc:t,className:n}=e,a=(0,l.useMemo)(()=>t&&t.items?t.items.flatMap(e=>{var t;return[e.url,null==e?void 0:null===(t=e.items)||void 0===t?void 0:t.map(e=>e.url)]}).flat().filter(Boolean).map(e=>null==e?void 0:e.split("#")[1]):[],[t]),s=function(e){let[t,n]=(0,l.useState)(void 0);return(0,l.useEffect)(()=>{let t=new IntersectionObserver(e=>{e.forEach(e=>{e.isIntersecting&&n(e.target.id||"")})},{rootMargin:"0% 0% -80% 0%"});return null==e||e.forEach(e=>{let n=document.getElementById(e);n&&t.observe(n)}),()=>{null==e||e.forEach(e=>{let n=document.getElementById(e);n&&t.unobserve(n)})}},[e]),t}(a);return(null==t?void 0:t.items)?(0,r.jsx)("aside",{className:(0,c.cn)("text-xs transition-transform xl:text-sm",n),children:(0,r.jsxs)("div",{className:(0,c.cn)("sticky top-16 -mt-10 max-h-[calc(var(--vh)-4rem)] space-y-2 overflow-y-auto pr-2 pt-16"),children:[(0,r.jsx)("p",{className:"text-sm font-medium uppercase",children:"On This Page"}),(0,r.jsxs)("span",{className:"child:w-auto mb-4 inline-flex flex-col space-y-2",children:[(0,r.jsxs)(o(),{href:"/docs",className:"group/link relative inline-flex items-center space-x-2 pb-1.5 uppercase text-neutral-600 dark:text-neutral-400",children:[(0,r.jsx)(f.poN,{className:"h-3 w-auto"}),(0,r.jsx)("span",{children:"Back to Docs"}),(0,r.jsx)("span",{className:"absolute inset-x-0 bottom-1 h-0.5 w-0 bg-current transition-width group-hover/link:w-full","aria-hidden":"true"})]}),(0,r.jsx)(d,{})]}),(0,r.jsx)(p,{tree:t,activeItem:s})]})}):null}function p(e){var t;let{tree:n,level:l=1,activeItem:a}=e;return(null==n?void 0:null===(t=n.items)||void 0===t?void 0:t.length)&&l<4?(0,r.jsx)("ul",{className:(0,c.cn)("m-0 list-none",{"pl-2":1!==l}),style:{paddingLeft:"".concat(.5*l,"rem")},children:n.items.map((e,t)=>{var n;return(0,r.jsxs)("li",{className:(0,i.Z)("mt-0 pt-2"),children:[(0,r.jsx)("a",{href:e.url,className:(0,c.cn)("inline-block text-sm no-underline",e.url==="#".concat(a)?"font-medium text-rose-600 dark:text-orange-300":"text-neutral-700 hover:text-neutral-900 dark:text-neutral-400"),children:e.title}),(null===(n=e.items)||void 0===n?void 0:n.length)?(0,r.jsx)(p,{tree:e,level:l+1,activeItem:a}):null]},t)})}):null}},8301:function(e,t,n){"use strict";n.r(t),n.d(t,{ScrollArea:function(){return s},ScrollBar:function(){return i}});var r=n(5001),l=n(8471),a=n(150),o=n(7190);let s=l.forwardRef((e,t)=>{let{className:n,children:l,...s}=e;return(0,r.jsxs)(a.fC,{ref:t,className:(0,o.cn)("relative overflow-hidden",n),...s,children:[(0,r.jsx)(a.l_,{className:"h-full w-full rounded-[inherit]",children:l}),(0,r.jsx)(i,{}),(0,r.jsx)(a.Ns,{})]})});s.displayName=a.fC.displayName;let i=l.forwardRef((e,t)=>{let{className:n,orientation:l="vertical",...s}=e;return(0,r.jsx)(a.gb,{ref:t,orientation:l,className:(0,o.cn)("flex touch-none select-none transition-colors","vertical"===l&&"h-full w-2.5 border-l border-l-transparent p-[1px]","horizontal"===l&&"h-2.5 border-t border-t-transparent p-[1px]",n),...s,children:(0,r.jsx)(a.q4,{className:"bg-border relative flex-1 rounded-full"})})});i.displayName=a.gb.displayName},7190:function(e,t,n){"use strict";n.d(t,{cn:function(){return a}});var r=n(9754),l=n(1171);function a(){for(var e=arguments.length,t=Array(e),n=0;n<e;n++)t[n]=arguments[n];return(0,l.m)((0,r.W)(...t))}},5552:function(){}},function(e){e.O(0,[354,831,327,278,373,335,237,744],function(){return e(e.s=2685)}),_N_E=e.O()}]);