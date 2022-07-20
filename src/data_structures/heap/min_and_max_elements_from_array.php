<?php

declare(strict_types=1);

namespace Ehunov\Interview\data_structures\heap;

final class min_and_max_elements_from_array
{
    public static function min(array $haystack, int $count): array
    {
        return self::passThroughHeap(new \SplMaxHeap(), $haystack, $count);
    }

    public static function max(array $haystack, int $count): array
    {
        return self::passThroughHeap(new \SplMinHeap(), $haystack, $count);
    }

    private static function passThroughHeap(\SplHeap $heap, array $haystack, int $count): array
    {
        for ($i = 0; $i < $count; $i++) {
            $heap->insert($haystack[$i]);
        }

        for ($iMax = count($haystack); $i < $iMax; $i++) {
            $heap->insert($haystack[$i]);
            $heap->extract();
        }

        return array_values(iterator_to_array($heap));
    }
}
