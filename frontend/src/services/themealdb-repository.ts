import { IMealDBRecipe, IMealDBRecipeAPIReponse } from "../types/themealdb";
import { IRecipe } from "../types/viewModels";

const API_URL = "https://www.themealdb.com/api/json/v1/1";

function mapData(recipeData: IMealDBRecipe): IRecipe {
  return {
    id: recipeData.idMeal,
    name: recipeData.strMeal,
    thumbnail: recipeData.strMealThumb,
    description: recipeData.strInstructions,
    ingredients: Object.keys(recipeData)
      .map((key) => {
        if (!key.startsWith("strIngredient")) {
          return { name: "", amount: "" };
        }
        const index = key.replace("strIngredient", "");
        return {
          name: recipeData[key] as string,
          amount: recipeData[`strMeasure${index}`] as string,
        };
      })
      .filter((item) => item.name),
    type: "unknown",
    difficulty: "unknown",
    servings: 0,
  };
}

export async function recipe(id: number): Promise<IRecipe | null> {
  const apiResponse = await fetch(`${API_URL}/lookup.php?i=${id}`);

  const responseData: IMealDBRecipeAPIReponse = await apiResponse.json();
  const recipeData = responseData.meals[0];

  return mapData(recipeData);
}

export async function search(query: string): Promise<IRecipe[]> {
  const apiResponse = await fetch(`${API_URL}/search.php?s=${query}`);
  const responseData: IMealDBRecipeAPIReponse = await apiResponse.json();

  const result: IRecipe[] = [];
  if (responseData.meals) {
    responseData.meals.forEach((element) => {
      result.push(mapData(element));
    });
  }
  return result;
}
