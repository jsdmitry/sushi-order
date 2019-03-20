import * as types from './actionsType';

export function getMenuFromCategorySuccess(data:any) {
  return {
    type: types.GET_MENU_FROM_CATEGORY_SUCCESS,
    data: data
  };
}
