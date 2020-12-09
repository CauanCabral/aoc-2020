<?php
class Bag
{
    public array $parents;

    public array $children;

    public function __construct(public string $name)
    {
        $this->parents = [];
        $this->children = [];
    }

    public function addChild(Bag $child, bool $reverseLink = true, ?int $limit = null): void
    {
        if (isset($this->parents[$child->name])) {
            throw new RuntimeException("Ciclyc dependency detected.");
        }

        if (!isset($this->children[$child->name])) {
            $this->children[$child->name] = [
                'limit' => $limit,
                'instance' => $child
            ];
        }

        if ($reverseLink) {
            $child->addParent(parent: $this, reverseLink: false);
        }
    }

    public function addParent(Bag $parent, bool $reverseLink = true): void
    {
        if (isset($this->children[$parent->name])) {
            throw new RuntimeException("Ciclyc dependency detected.");
        }

        if (!isset($this->children[$parent->name])) {
            $this->parents[$parent->name] = $parent;
        }

        if ($reverseLink) {
            $parent->addChild(child: $this, reverseLink: false);
        }
    }

    public function eventuallyParents(): array
    {
        $eventually = $this->parents;
        foreach ($eventually as $pName => $parent) {
            $eventually += $parent->eventuallyParents();
        }

        return $eventually;
    }

    public function mustContain(): int
    {
        $contain = 0;
        foreach ($this->children as $child) {
            $contain += $child['limit'] + ($child['limit'] * $child['instance']->mustContain());
        }

        return $contain;
    }
}