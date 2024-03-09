-- Definición de un tipo de dato para representar a un estudiante
data Estudiante = Estudiante { nombre :: String, nota :: Float } deriving (Show)

-- Función para calcular la lista de mejores estudiantes
mejoresEstudiantes :: Float -> [Estudiante] -> [Estudiante]
mejoresEstudiantes umbral = filter ((>= umbral) . nota)

-- Función para calcular el promedio general del curso
promedioGeneral :: [Estudiante] -> Float
promedioGeneral estudiantes = sum (map nota estudiantes) / fromIntegral (length estudiantes)

-- Función principal para probar las funciones
main :: IO ()
main = do
    let estudiantes = [ Estudiante "Leiner" 70.0, Estudiante "Steven" 100.0, Estudiante "Walter" 85.0, Estudiante "Karina" 50.0]
    putStrLn "Lista de estudiantes:"
    print estudiantes
    putStrLn "Mejores estudiantes con nota mayor o igual a 15.0:"
    print $ mejoresEstudiantes 85.0 estudiantes
    putStrLn $ "Promedio general del curso: " ++ show (promedioGeneral estudiantes)
