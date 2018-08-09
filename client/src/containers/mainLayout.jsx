import React, {Component} from 'react';
import PersistentDrawer from "../components/persistentDrawer"
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import {selectCategory} from "../actions/categoryActions";

class MainLayout extends Component {
  render() {
    return (
      <PersistentDrawer selectCategory={this.props.selectCategory} selectedCategoryID={this.props.selectedCategoryID}/>
    );
  }
}

function mapStateToProps(state) {
  return {
    selectedCategoryID: state.categories.selectedCategoryID
  }
}

function mapDispatchToProps(dispatch) {
  return {
    selectCategory: bindActionCreators(selectCategory, dispatch)
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(MainLayout);
