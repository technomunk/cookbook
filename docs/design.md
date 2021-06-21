# Cookbook design

The core principle of the cookbook is to provide a convenient recipe store. Convenience means powerful, but easy queries and presenting results in human-readable way.

## Terminology

- **Ingredient**: a product with a name, ex: *egg*, *pizza*, *bacon & egg*.
- **Rate**: proportional amount of an ingredient. Can be expressed in *units* (ex: 1 egg, 2 hamburger buns), *weight* (200g of white rice) or *volume* (200ml of water).
- **Recipe**: a notion of a process that takes ingredient(s) at a known rate and converts them to something else.
- **Recipe process**: a process is action performed to a set of ingredients to transform them in some way, ex: *bake*, *cook*, *slice* and *mix*. Some processes may require the recipe to specify a particular attribute, ex: *baking* requires *temperature* and *duration*.
