import { combineReducers } from 'redux';
import categories from './categories';
import menu from './menu';

export default combineReducers({
  categories,
  menu
});

