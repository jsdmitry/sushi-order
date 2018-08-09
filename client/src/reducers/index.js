import { combineReducers } from 'redux';
import categories from '../reducers/categories';
import menu from '../reducers/menu';

export default combineReducers({
  categories,
  menu
});

