<?php
//15.三数之和



function threeSum($nums) {
    $res = array();
    for ($i=0;$i<count($nums)-2;$i++){
        for ($j=$i+1;$j<count($nums)-1;$j++){
            for ($k=$j+1;$k<count($nums);$k++){
                if($nums[$i]+$nums[$j]+$nums[$k] == 0){
                    $temp = [$nums[$i],$nums[$j],$nums[$k]];
                    sort($temp);
                    if (!in_array($temp,$res)){
                        array_push($res,$temp);
                    }
                }
            }
        }
    }
    return $res;
}

function threeSum1($nums) {
    $res = array();
    if (count($nums)<3){
        return $res;
    }
    return $res;
}

$nums = [-1,0,1,2,-1,-4];
var_dump(threeSum($nums));



