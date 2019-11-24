/**
 * Returns the base url of the plugin
 * installation.
 *
 * @return {string} instance base URL
 */
function getBaseURL() {
    const url = new URL(window.location.href);
    return `${url.protocol}//${url.host}`;
}

function popupCenter(url, title, w, h) {
    // Fixes dual-screen position                            Most browsers       Firefox
    const dualScreenLeft = window.screenLeft === undefined ? window.screenX : window.screenLeft; // eslint-disable-line no-undefined
    const dualScreenTop = window.screenTop === undefined ? window.screenY : window.screenTop; // eslint-disable-line no-undefined

    let width;
    let height;

    if (window.innerWidth) {
        width = window.innerWidth;
    } else {
        width = document.documentElement.clientWidth ? document.documentElement.clientWidth : screen.width;
    }

    if (window.innerHeight) {
        height = window.innerHeight;
    } else {
        height = document.documentElement.clientHeight ? document.documentElement.clientHeight : screen.height;
    }

    const left = ((width / 2) - (w / 2)) + dualScreenLeft;
    const top = ((height / 2) - (h / 2)) + dualScreenTop;
    const newWindow = window.open(url, title, 'scrollbars=yes, width=' + w + ', height=' + h + ', top=' + top + ', left=' + left);

    if (newWindow != null) {
        newWindow.focus();
    }

    return newWindow;
}

var lut = [];
for (var i = 0; i < 256; i++) {
    lut[i] = (i < 16 ? '0' : '') + (i).toString(16);
}

function uuidv4() {
    var d0 = Math.random() * 0xffffffff | 0;
    var d1 = Math.random() * 0xffffffff | 0;
    var d2 = Math.random() * 0xffffffff | 0;
    var d3 = Math.random() * 0xffffffff | 0;
    return lut[d0 & 0xff] + lut[d0 >> 8 & 0xff] + lut[d0 >> 16 & 0xff] + lut[d0 >> 24 & 0xff] + '-' +
        lut[d1 & 0xff] + lut[d1 >> 8 & 0xff] + '-' + lut[d1 >> 16 & 0x0f | 0x40] + lut[d1 >> 24 & 0xff] + '-' +
        lut[d2 & 0x3f | 0x80] + lut[d2 >> 8 & 0xff] + '-' + lut[d2 >> 16 & 0xff] + lut[d2 >> 24 & 0xff] +
        lut[d3 & 0xff] + lut[d3 >> 8 & 0xff] + lut[d3 >> 16 & 0xff] + lut[d3 >> 24 & 0xff];
}

export default {
    getBaseURL,
    popupCenter,
    uuidv4,
};
