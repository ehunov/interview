<?php

namespace Ehunov\Interview\data_structures\heap;

use PHPUnit\Framework\TestCase;

final class min_and_max_elements_from_array_test extends TestCase
{
    public function testMin(): void
    {
        $randomNumbersArray = $this->generate($randomElementCount = 100_000, $maxAndDefaultValue = 10, 1_000_000);

        $minArray = min_and_max_elements_from_array::min($randomNumbersArray, $randomElementCount);

        sort($minArray);
        for ($i = 0; $i < $randomElementCount; $i++) {
            $this->assertLessThan($maxAndDefaultValue, $minArray[$i]);
        }
    }

    public function testMax(): void
    {
        $randomNumbersArray = $this->generate($randomElementCount = 100_000, $maxAndDefaultValue = 10, 1_000_000);

        $maxArray = min_and_max_elements_from_array::max($randomNumbersArray, $randomElementCount);

        for ($i = 0; $i < $randomElementCount; $i++) {
            $this->assertEquals($maxAndDefaultValue, $maxArray[$i]);
        }
    }

    private function generate(int $randomElementsCount, int $maxAndDefaultValue, int $count): array
    {
        $array = array_fill(0, $count, $maxAndDefaultValue);

        $existedRandomKeys = [];
        for ($i = 0; $i < $randomElementsCount; $i++) {
            $array[$this->randomButNotExisted($count-1, $existedRandomKeys)] = random_int(0, $maxAndDefaultValue-1);
        }

        return $array;
    }

    private function randomButNotExisted(int $max, array &$existed): int
    {
        $rand = random_int(0, $max);

        if (!in_array($rand, $existed, true)) {
            $existed[] = $rand;
            return $rand;
        }

        return $this->randomButNotExisted($max, $existed);
    }
}
