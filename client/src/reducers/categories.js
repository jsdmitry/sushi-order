import * as types from '../actions/actionsType';

const initialState = {
  data: []
};

export default function(state = initialState, action) {
  switch(action.type) {
    case types.CATEGORIES_LIST_SUCCESS:
      return { ...state, data: action.data };
    default:
      return state;
  }
}
