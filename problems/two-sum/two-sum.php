<?php
function twoSum($nums, $target) {
	$indexMap = [];
	$output = [];
	foreach ($nums as $index=>$value) {
		$otherNumberKey = (string)($target - $value);
		if (isset($indexMap[$otherNumberKey])) {
			$output = [$indexMap[$otherNumberKey], $index];
			break;
		}
		$indexMap[(string)$value] = $index;
	}
	return $output;
}
var_dump(twoSum([2,7,11,15],9));