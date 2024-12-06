namespace day2;
using System.IO;

public static class PartOne
{
    public static void Run()
    {
        IEnumerable<string> reports = File.ReadAllLines("input.txt");
        
        // delete duplicates
        reports = reports.Where(d => d.Split(" ").Distinct().Count() == d.Split(" ").Length);
        
        // delete not sorted 
        reports = reports.Where(d =>
        {
            var levels = d.Split(" ");
            return levels.OrderBy(int.Parse).SequenceEqual(levels) ||
                   levels.OrderByDescending(int.Parse).SequenceEqual(levels);
        });

        // delete those that doesn't follow the rule (increment or decrement by 1 or 3)
        reports = reports.Where(d =>
        {
            var levels = d.Split(" ").Select(int.Parse).ToList();
            var incrementLevels =  levels.Zip(levels.Skip(1), (a, b) => b - a).All(x => x is >= 1 and <= 3);
            var decrementLevels = levels.Zip(levels.Skip(1), (a, b) => a - b).All(x => x is >= 1 and <= 3);
            return incrementLevels || decrementLevels;
        });
        
        Console.WriteLine(reports.Count());
    }
}
