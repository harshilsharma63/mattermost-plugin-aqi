import {connect} from 'react-redux';

import Selectors from '../../selectors';

import LeftSidebarHeader from './left_sidebar_header';

const mapStateToProps = (state) => ({
    pollutionData: Selectors.getPollutionData(state),
});

export default connect(mapStateToProps)(LeftSidebarHeader);
