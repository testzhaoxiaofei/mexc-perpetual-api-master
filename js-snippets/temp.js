// https://static.mocortech.com/futures-v3/_next/static/chunks/app/%5Blocale%5D/layout-deef8df03d50baa8.js

async function submitOrder(e) {
    var t;
    let r = null === (t = JSON.parse(window.localStorage.getItem("mxc.settings"))) || void 0 === t ? void 0 : t.triggerProtect
      , o = (0,
    s.Li)()
      , i = "".concat(m, "/private/order/create");
    return (null == o ? void 0 : o.mhash) && (i += "?mhash=".concat(o.mhash)),
    logMarketOrderInfo({
        params: e
    }),
    (0,
    n.ZP)(i, {
        method: "POST",
        body: {
            ...e,
            priceProtect: r,
            ...o
        },
        needLogin: !0
    }).then(t => {
        if (null == t ? void 0 : t.success) {
            var r, n;
            null === (n = window) || void 0 === n || null === (r = n.dataLayer) || void 0 === r || r.push({
                event: "furture_trade"
            }),
            logMarketOrderInfo({
                params: e,
                res: t
            })
        }
        return t
    }
    )
}

// submitOrder 的s.Li函数就是getFingerPrintData
getFingerPrintData = () => {
    try {
        var e, t, r, o, i, a;
        let l = (null === (e = window) || void 0 === e ? void 0 : e.store.getState()) || {}
          , s = l.auth
          , c = null !== (a = null === (i = window) || void 0 === i ? void 0 : null === (o = i.fp) || void 0 === o ? void 0 : null === (r = o.getFpDataSync) || void 0 === r ? void 0 : r.call(o, {
            userId: null == s ? void 0 : null === (t = s.loginMember) || void 0 === t ? void 0 : t.memberId,
            scene: n.ZWA.DM_ORDER
        })) && void 0 !== a ? a : void 0;
        if ("object" != typeof c)
            return {};
        return c
    } catch (e) {
        return {}
    }
}

function pr(e) {
    var r = e.scene
      , n = e.userId
      , i = void 0 === n ? "" : n
      , o = i ? {
        member_id: i
    } : {}
      , a = function(t) {
        void 0 === t && (t = 0);
        var e = nr(t)
          , r = e[0];
        return !e[1] && r ? r : (ur(t),
        ir())
    }(r)
      , s = Ke() || {};
    return hr(a, t(t({}, s), o))
}

// de = "1b8c71b668084dda9dc0285171ccf753"

function nr(t) {
    var e = function(t) {
        var e = window.localStorage.getItem(ke) || "{}";
        return e ? JSON.parse(e)[t] : null
    }(t);
    if (!e)
        return [e, !0];
    var r = e.expire
      , n = e.config
      , i = e.overdue
      , o = JSON.parse(Ve(n, de));
    return 0 === i ? [o, !1] : [o, ot() > r]
}

// qe = "aes-256-gcm"
// , We = "utf8";
function Ve(t, e) {
    var r = Ue.from(e, We)
      , n = Ue.from(t, "base64")
      , i = n.slice(0, 12)
      , o = n.slice(12, -16)
      , a = n.slice(-16)
      , s = Fe.G_(qe, r, i);
    return s.setAuthTag(a),
    s.update(o, void 0, We) + s.final(We)
}

// createDecipheriv
function d(t, e, r) {
    var s = o[t.toLowerCase()];
    if (!s)
        throw new TypeError("invalid suite type");
    if ("string" == typeof r && (r = i.from(r)),
    "GCM" !== s.mode && r.length !== s.iv)
        throw new TypeError("invalid iv length " + r.length);
    if ("string" == typeof e && (e = i.from(e)),
    e.length !== s.key / 8)
        throw new TypeError("invalid key length " + e.length);
    return "stream" === s.type ? new a(s.module,e,r,!0) : "auth" === s.type ? new n(s.module,e,r,!0) : new c(s.module,e,r)
}

// Re = "mexc_local_fingerprint_sys_info"
// mexc_local_fingerprint_sys_info的键存储浏览器指纹信息
function Ke() {
    var t = window.localStorage.getItem(Re);
    return t ? JSON.parse(t) : null
}

function hr(t, e) {
    var r, n, i = ge(), o = e.mtoken, a = e.mhash, s = t.chash, u = t.parameters, f = t.data_upload, c = Fe.O6(16).toString("hex"),
        h = (r = c,
            n = i,
            Fe.r0({
                key: n,
                padding: Fe._G.RSA_PKCS1_PADDING
            }, Ue.from(r)).toString("base64")),
        d = (new Date).getTime(), l = function (t, e) {
            for (var r = {}, n = 0, i = e; n < i.length; n++) {
                var o = i[n];
                t.hasOwnProperty(o) && (r[o] = t[o])
            }
            return r
        }(e, 1 === f ? u : ["mtoken"]);
    return {
        p0: ze(JSON.stringify(l), c),
        k0: h,
        chash: s,
        mtoken: o,
        ts: d,
        mhash: a
    }
}

// Ue = r(8764).lW -- createDecipherIv
// , qe = "aes-256-gcm"
// , We = "utf8";
function ze(t, e) {
    var r = Ue.from(e, We)
      , n = Fe.O6(12)
      , i = Fe.CW(qe, r, n)
      , o = i.update(t, We, "hex") + i.final("hex")
      , a = i.getAuthTag();
    return Ue.from(n.toString("hex") + o + a.toString("hex"), "hex").toString("base64")
}
