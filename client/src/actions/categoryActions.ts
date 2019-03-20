import * as types from './actionsType';

export function getCategoriesSuccess(data:any) {
  return {
    type: types.CATEGORIES_LIST_SUCCESS,
    data: data
  };
}

export function selectCategory(id:number) {
  return {
    type: types.SELECT_CATEGORY,
    id: id
  };
}
