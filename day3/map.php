<?php
class Map
{
    /**
     * Number of real lines on the pattern
     */
    private int $realLines;

    /**
     * Number of real columns on the pattern
     */
    private int $realColumns;

    public function __construct(
        protected array $pattern,
        protected ?int $width = null,
        protected ?int $height = null,
    )
    {
        $this->realLines = count($this->pattern);
        $this->realColumns = strlen($this->pattern[0]);

        if (!$this->width) {
            $this->width = $this->realColumns;
        }

        if (!$this->height) {
            $this->height = $this->realLines;
        }
    }

    public function setWidth(int $width): self
    {
        $this->width = $width;

        return $this;
    }

    public function getWidth(): int
    {
        return $this->width;
    }

    public function setHeight(int $height): self
    {
        $this->height = $height;

        return $this;
    }

    public function getHeight(): int
    {
        return $this->height;
    }

    public function get(int $x, int $y): int|string
    {
        if ($x >= $this->width) {
            throw new RuntimeException("Invalid column position");
        }

        if ($y >= $this->height) {
            throw new RuntimeException("Invalid line position");
        }

        $a = $x % $this->realColumns;
        $b = $y % $this->realLines;

        return $this->pattern[$b][$a];
    }

    /**
     * Draw a $width x $height map based on pattern
     *
     * @return void
     */
    public function draw(): void
    {
        for ($i = 0; $i < $this->height; $i++) {
            $line = $this->pattern[$i % $this->realLines];
            echo str_pad($line, $this->width, $line) . "\n";
        }
    }
}