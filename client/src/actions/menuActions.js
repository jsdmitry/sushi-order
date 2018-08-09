import * as types from '../actions/actionsType';

export function getMenuFromCategorySuccess(data) {
  return {
    type: types.GET_MENU_FROM_CATEGORY_SUCCESS,
    data: data
  };
}
