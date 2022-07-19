<?php

// https://leetcode.com/problems/merge-sorted-array/

namespace Ehunov\Interview\leetcode\data_structures\array;

final class merge_sorted
{
    public static function merge(array &$nums1, int $m, array $nums2, int $n): array
    {
        $i = $m + $n - 1;
        $m--;
        $n--;

        while ($n >= 0) {
            if ($m >= 0 && $nums1[$m] >= $nums2[$n]) {
                $nums1[$i] = $nums1[$m];
                $m--;
            } else {
                $nums1[$i] = $nums2[$n];
                $n--;
            }
            $i--;
        }

        return $nums1;
    }
}
