<?php

namespace Ehunov\Interview\data_structures\heap;

use PHPUnit\Framework\TestCase;

class top_k_frequent_test extends TestCase
{
    public function testTop(): void
    {
        $this->assertEquals([1, 2], top_k_frequent::top([1,1,1,2,2,3], 2));
    }
}
