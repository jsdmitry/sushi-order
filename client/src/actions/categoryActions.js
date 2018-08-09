import * as types from '../actions/actionsType';

export function getCategoriesSuccess(data) {
  return {
    type: types.CATEGORIES_LIST_SUCCESS,
    data: data
  };
}

export function selectCategory(id) {
  return {
    type: types.SELECT_CATEGORY,
    id: id
  };
}
