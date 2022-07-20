<?php declare(strict_types=1);

namespace Ehunov\Interview\data_structures\array;

use PHPUnit\Framework\TestCase;

final class merge_sorted_test extends TestCase
{
    /**
     * @dataProvider provideTestMergeData
     */
    public function testMerge(array $nums1, int $n, array $nums2, int $m, array $expected): void
    {
        $this->assertEquals($expected, merge_sorted::merge($nums1, $n, $nums2, $m));
    }

    public function provideTestMergeData(): iterable
    {
        yield [
            [1, 2, 3, 0, 0, 0],
            3,
            [2, 5, 6],
            3,
            [1, 2, 2, 3, 5, 6],
        ];
        yield [
            [1],
            1,
            [],
            0,
            [1],
        ];
        yield [
            [0],
            0,
            [1],
            1,
            [1],
        ];
        yield [
            [2, 0],
            1,
            [1],
            1,
            [1, 2],
        ];
    }
}
