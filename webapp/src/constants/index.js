import Utils from '../utils';

import {PLUGIN_NAME} from './manifest';
import SVGS from './svgs';

//
// Define our URL constants
//
const URL_BASE = `${Utils.getBaseURL()}/plugins/${PLUGIN_NAME}`;

const URL_REFRESH_DATA = `${URL_BASE}/refresh`;

const ACTIONS = {
    RECEIVE_POLLUTION_DATA: `${PLUGIN_NAME}_receive_pollution_data`,
};

//
// Export the constants
//
export default {
    PLUGIN_NAME,
    URL_BASE,
    URL_REFRESH_DATA,
    SVGS,
    ACTIONS,
};
