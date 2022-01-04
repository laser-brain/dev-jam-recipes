import { IRecipe } from "../types/viewModels";

interface IRecipesAPIResponse {
  recipes: IRecipe[];
}

const data: IRecipe[] = [];

async function loadData() {
  var apiResponse = await fetch("http://localhost:8080/recipes?page=1&size=20");
  var responseData: IRecipesAPIResponse = await apiResponse.json();

  for (let item of responseData.recipes) {
    data.push(item);
  }
}

export async function recipes(): Promise<IRecipe[]> {
  if (data.length === 0) {
    await loadData();
  }
  return data;
}

export async function recipe(id: number): Promise<IRecipe | null> {
  if (data.length === 0) {
    await loadData();
  }

  const filtered = data.filter((recipe) => recipe.id === id);
  if (filtered.length === 1) {
    return filtered[0];
  }

  return null;
}

export async function search(query: string) {
  if (data.length === 0) {
    await loadData();
  }

  return data.filter((recipe) => recipe.name.toLowerCase().indexOf(query.toLowerCase()) !== -1);
}
