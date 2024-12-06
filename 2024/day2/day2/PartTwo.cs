namespace day2;

public static class PartTwo
{
    public static void Run()
    {
        IEnumerable<string> reports = File.ReadAllLines("input.txt");
        
        // Remove those with 2 or more duplicates (there is any duplicated but is a good practice to remove them as an edge case)
        reports = reports.Where(d => d.Split(" ").GroupBy(x => x).Any(g => g.Count() <= 2));

        var total = 0;
        foreach (var report in reports)
        {
            // Checking if it follows the rules without any modification
            if (FollowRules(report))
            {
                total += 1;
                continue;
            }
            
            // Validating if removing an element from the list makes it follow the rules
            var levelsQuantity = report.Split(" ").Count();
            for (var i = 0; i < levelsQuantity; i++)
            {
                var levels = report.Split(" ").Select(int.Parse).ToList();
                levels.RemoveAt(i);
                var auxReport = string.Join(" ", levels);
                if (FollowRules(auxReport))
                {
                    total += 1;
                    break;
                }
            }
        }
        
        Console.WriteLine(total);
    }

    private static bool FollowRules(string report)
    {
        var levels = report.Split(" ").Select(int.Parse).ToList();
        var isOrdered = levels.OrderBy(v => v).SequenceEqual(levels);
        var isOrderedDesc = levels.OrderByDescending(v => v).SequenceEqual(levels);
        var incrementLevels =  levels.Zip(levels.Skip(1), (a, b) => b - a).All(x => x is >= 1 and <= 3);
        var decrementLevels = levels.Zip(levels.Skip(1), (a, b) => a - b).All(x => x is >= 1 and <= 3);
        var followRules = incrementLevels || decrementLevels;
        return (isOrdered || isOrderedDesc) && followRules;
    }
}