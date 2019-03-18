import * as types from '../actions/actionsType';

const initialState = {
  data: []
};

export default function(state = initialState, action) {
  switch(action.type) {
    case types.CATEGORIES_LIST_SUCCESS:
      return { ...state, data: action.data };
    case types.SELECT_CATEGORY:
      return { ...state, selectedCategoryID: action.id };
    default:
      return state;
  }
}
