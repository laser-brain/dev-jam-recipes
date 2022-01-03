import { IRecipe } from "../types/viewModels";

const data: IRecipe[] = [];
  for (let index = 1; index <= 20; index++) {
    data.push({
      id: index,
      name: `recipe ${index}`,
      thumbnail: "https://www.themealdb.com/images/logo-small.png",
      description: `This is the description of the recipe #${index}`
    });
  }

export function recipes(): IRecipe[] {  
  return data;
}

export function recipe(id: number) : IRecipe | null {
  const filtered = data.filter(recipe => recipe.id === id);
  if(filtered.length === 1){
    return filtered[0];
  }

  return null;
}
