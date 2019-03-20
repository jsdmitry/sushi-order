import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import Menu from '../components/menu';
import {getMenuFromCategorySuccess} from '../actions/menuActions';
import axios from 'axios';

class MenuContainer extends Component {
  componentWillUpdate(nextProps) {
    if(nextProps.selectedCategoryID !== this.props.selectedCategoryID) {
      axios
      .get(`http://localhost:8080/sushi-data/menu/category/` + nextProps.selectedCategoryID, {
        crossDomain: true
      })
      .then(responce => {
        this.props.getMenuFromCategorySuccess(responce.data.data);
      });
    }
  }

  render() {
    return (
      <Menu data={this.props.menu.data} colCount={3}/>
    );
  }
}

function mapStateToProps(state) {
  return {
    menu: state.menu
  }
}

function mapDispatchToProps(dispatch) {
  return {
    getMenuFromCategorySuccess: bindActionCreators(getMenuFromCategorySuccess, dispatch)
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(MenuContainer);
