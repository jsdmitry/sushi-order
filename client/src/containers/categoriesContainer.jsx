import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import CategoriesList from '../components/categoriesList';
import {getCategoriesSuccess} from '../actions/categoryActions';
import axios from 'axios';

class CategoriesContainer extends Component {
  componentDidMount() {
    axios
    .get(`http://localhost:8080/sushi-data/categories/`, {
      crossDomain: true
    })
    .then(responce => {
      this.props.getCategoriesSuccess(responce.data.data);
    })
  }

  render() {
    return (
      <CategoriesList categories={this.props.categories} onCategorySelectionChanged={this.props.onCategorySelectionChanged}/>
    );
  }
}

function mapStateToProps(state) {
  return {
    categories: state.categories
  }
}

function mapDispatchToProps(dispatch) {
  return {
    getCategoriesSuccess: bindActionCreators(getCategoriesSuccess, dispatch)
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(CategoriesContainer);
