import Constants from '../constants';

export const receivePollutionData = (pollutionData) => (dispatch) => {
    dispatch({
        type: Constants.ACTIONS.RECEIVE_POLLUTION_DATA,
        pollutionData,
    });
};
