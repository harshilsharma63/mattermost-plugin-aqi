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
export default {
    getBaseURL,
};
