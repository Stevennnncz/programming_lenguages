-- Definición de un tipo algebraico para representar productos
data Producto = Laptop { modelo :: String, precio :: Float }
              | TelefonoMovil { marca :: String, precio :: Float }
              deriving (Show)

-- Lista de productos
productos :: [Producto]
productos = [ Laptop "HP Pavilion" 800.0
            , TelefonoMovil "Samsung Galaxy" 500.0
            , Laptop "Dell XPS" 1200.0
            , TelefonoMovil "iPhone 13" 1000.0
            ]

-- Calcular el precio total de los productos
precioTotal :: [Producto] -> Float
precioTotal = sum . map precio

main :: IO ()
main = do
    putStrLn "Lista de productos:"
    print productos
    putStrLn $ "Precio total de los productos: " ++ show (precioTotal productos)