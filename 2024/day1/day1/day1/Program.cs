void PartOne()
{
    var total = 0;
    var lines = File.ReadAllLines("input.txt");
    var firstList = new List<int>();
    var secondList = new List<int>();
    foreach (var line in lines)
    {
        var lineData = line.Split("   ");
        firstList.Add(int.Parse(lineData.First()));
        secondList.Add(int.Parse(lineData.Last()));
    }

    firstList.Sort();
    secondList.Sort();
    for (var i = 0; i < firstList.Count; i++)
    {
        total += Math.Abs(firstList[i] - secondList[i]);
    }

    Console.WriteLine(total);
}

void PartTwo()
{
    var lines = File.ReadAllLines("input.txt");
    var firstList = new List<int>();
    var secondList = new List<int>();
    foreach (var line in lines)
    {
        var lineData = line.Split("   ");
        firstList.Add(int.Parse(lineData.First()));
        secondList.Add(int.Parse(lineData.Last()));
    }
    var total = firstList.Sum(value => value * secondList.Count(v => v == value));
    Console.WriteLine(total);
}

PartOne();
PartTwo();