// https://github.com/nimojs/learn/leetcode/problems/can-place-flowers/
class Solution {
    function canPlaceFlowers($flowerbed, $n) {
      array_push($flowerbed, 0);
      array_unshift($flowerbed, 0);
      $currentRangeEmptyCount = 0;
      $canPlaceCount = 0;
      $flowerbedLen = count($flowerbed);
      foreach($flowerbed as $index=>$item) {
        if ($item === 0) {
           $currentRangeEmptyCount++;
        }
        if ($index === $flowerbedLen-1 || $item === 1) {
          if ($currentRangeEmptyCount >= 3) {
            $canPlaceCount = $canPlaceCount + floor($currentRangeEmptyCount-1)/2;
          }
        }
        if ($item === 1){
          $currentRangeEmptyCount = 0;
        }
      }
      return $canPlaceCount >= $n;
  }
}