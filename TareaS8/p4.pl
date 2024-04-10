% Facts: Define the available dishes categorized by type
dish_type(entradas, [guacamole, ensalada, consome, tostadas_caprese]).
dish_type(carne, [filete_de_cerdo, pollo_al_horno, carne_en_salsa]).
dish_type(pescado, [tilapia, salmon, trucha]).
dish_type(postre, [flan, nueces_con_miel, naranja_confitada, flan_de_coco]).

% Rules: Define the predicate to check if a dish belongs to a certain type
belongs_to(Dish, Type) :-
    dish_type(Type, Dishes), % Get the list of dishes for the given type
    member(Dish, Dishes). % Check if the dish belongs to the list of dishes for that type

% Example usage:
% ?- belongs_to(guacamole, entradas). => true
% ?- belongs_to(pollo_al_horno, carne). => true
% ?- belongs_to(salmon, entradas). => false
