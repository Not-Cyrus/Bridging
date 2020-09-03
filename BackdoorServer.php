<?php
    $LoginArray = [
        "user1" => "password1",
    ];
    if ($_SERVER['REQUEST_METHOD'] == "POST"){
        if (isset($_GET['Username']) && isset($_GET['Password']) && isset($_GET['Command'])){
            $Username = $_GET['Username'];
            $Password = $_GET['Password'];
            $Command = $_GET['Command'];
            if (array_key_exists($Username,$LoginArray) && $LoginArray[$Username] == $Password){
                $File = file_get_contents("Commands.json");
                $Array = json_decode($File);
                $Player = explode(" ",$Command);
                $Array[$Player[1]] = $Command;
                $Data = json_encode($Array);
                file_put_contents("Commands.json",$Data);
                sleep(15);
                array_pop($Array);
                $Data = json_encode($Array);
                file_put_contents("Commands.json",$Data);
            } 
        }
    }
?>