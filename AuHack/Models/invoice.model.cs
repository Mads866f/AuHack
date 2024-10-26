namespace AuHack.Models;

public class Invoice
{
    public DateTime Date { get; set; }
    public string InvoiceNumber { get; set; }
    public int Amount { get; set; }
    public string Description { get; set; }
    public string Sender { get; set; }
    public string Receiver { get; set; }
    public DateTime PaymentDate { get; set; }
}


