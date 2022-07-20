<?php

// https://leetcode.com/problems/top-k-frequent-words/

declare(strict_types=1);

namespace Ehunov\Interview\data_structures\heap;

final class top_k_frequent
{
    public static function top(array $nums, int $k): array
    {
        $freq = [];

        foreach ($nums as $num) {
            $freq[$num] = new Number(
                $num,
                ($freq[$num]?->frequency ?? 0) + 1
            );
        }

        $freq = array_values($freq);
        $minHeap = new NumberMinHeap();

        for ($i = 0; $i < $k; $i++) {
            $minHeap->insert($freq[$i]);
        }

        for ($iMax = count($freq); $i < $iMax; $i++) {
            $minHeap->insert($freq[$i]);
            $minHeap->extract();
        }

        return array_map(static fn(Number $number) => $number->value, iterator_to_array($minHeap));
    }
}

final class Number
{
    public function __construct(public int $value, public int $frequency)
    {
    }
}

final class NumberMinHeap extends \SplHeap {
    protected function compare(mixed $value1, mixed $value2): int
    {
        return $value2->frequency <=> $value1->frequency;
    }
}
